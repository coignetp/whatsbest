<template>
  <div class="mash-create">
    <router-link to="/mash/42/result" tag="button">Let's start</router-link> <br />
    <button @click="startTournament">Test start</button>
    <hr />
    <div class="mash-create-choices">
      <!-- <picture-input 
        v-for="(img, index) in images" :key="index" ref="chosen"
        @change="onChange(index)" 
        @remove="onRemove(index)"
        width="130" 
        height="200" 
        margin="16" 
        accept="image/jpeg,image/png,gif" 
        size="10" 
        :removable="true"
        :customStrings="{
          drag: 'Drag a choice'
        }">
      </picture-input>
      <hr /> -->
      <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"></vue-dropzone>
    </div>
  </div>
</template>

<script>
// import PictureInput from 'vue-picture-input';
import vue2Dropzone from "vue2-dropzone";
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import axios from 'axios';

const toBase64 = file => new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
});

export default {
  name: 'MashCreate',
  components: {
    // PictureInput,
    vueDropzone: vue2Dropzone,
  },
  data() {
    return {
      images: [],
      dropzoneOptions: {
        url: 'https://httpbin.org/post',
        resizeWidth: 200,
        resizeHeight: 300,
        maxFilesize: 0.5,
        addRemoveLinks: true,
        headers: { "My-Awesome-Header": "header value" },
      }
    }
  },
  created() {
    this.dropzoneOptions.accept = this.acceptUploadedFile;
    this.dropzoneOptions.removedfile = this.removeUploadedFile;
  },
  methods: {
    // onChange (index) {
    //   if (index === this.images.length - 1) {
    //     this.images.push(null);
    //   }
    //   if (index >= 0 && index < this.images.length) {
    //     this.images[index] = this.$refs.chosen[index].file;
    //   }
    //   console.log(this.images);
    // },
    // onRemove (index) {
    //   if (index >= 0 && index < this.images.length) {
    //     this.images.splice(index, 1);
    //     let i;
    //     for(i = this.images.length - 2; i >= index; i--) {
    //       console.log(i);
    //     }
    //   }
    //   console.log(this.images);
    // },
    async acceptUploadedFile(file, done) {
      this.images.push([file.upload.uuid, file]);
      done();
    },
    removeUploadedFile(file) {
      // console.log(file.upload.uuid);
      var index;
      for (index = 0; index < this.images.length; index++) {
        if (this.images[index][0] === file.upload.uuid) {
          this.images.splice(index, 1);
          return file.previewElement.parentNode.removeChild(file.previewElement);
        }
      }
      return false;
    },
    async createTournamentRequest() {
      let tournament = {
        "question": "What's best?",
        "size": this.images.length,
        "type": 1,
        "choices": [],
      };

      let img;
      for (img of this.images) {
        if (img != null) {
          tournament["choices"].push(
            await toBase64(img[1])
          );
        }
      }

      return tournament;
    },

    async startTournament() {
      console.log("In start Tournament");
      if (this.images) {
        const tournament = await this.createTournamentRequest()
  
        console.log(tournament);
        const config = {
          headers: {
            'content-type': 'application/json'
          }
        };

        try {
          const res = await axios.post("http://localhost:8081/create", tournament, config);
          if (res.status == 200) {
            this.image = '';
              console.log("Tournament successfully created");
              window.location.replace("#/mash/" + res.data["id"]);
          } else {
            console.error(res);
          }
        } catch (e) {
          console.error(e);
        }
      }
    }
  }
}

</script>

<style scoped>
#mash-create {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: crimson;
  margin-top: 60px;
}

h1, h2 {
  font-weight: normal;
}
</style>