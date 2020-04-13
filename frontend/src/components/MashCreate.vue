<template>
  <div class="mash-create">
    <button @click="startTournament">Let's start!</button>
    <hr />
    <div class="mash-create-choices">
      <b-container fluid>
        <b-row class="my-1">
          <b-col sm="3">
            <label>Question: </label>
          </b-col>
          <b-col sm="9">
            <b-form-input v-model="question" type="text" placeholder="What's best?"></b-form-input>
          </b-col>
        </b-row>
        <vue-dropzone ref="dropzone" id="dropzone" :options="dropzoneOptions"></vue-dropzone>
      </b-container>
    </div>
  </div>
</template>

<script>
// import PictureInput from 'vue-picture-input';
import vue2Dropzone from "vue2-dropzone";
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

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
    vueDropzone: vue2Dropzone,
  },
  data() {
    return {
      images: [],
      question: "",
      dropzoneOptions: {
        url: 'https://httpbin.org/post',
        resizeWidth: 200,
        resizeHeight: 300,
        maxFilesize: 0.5,
        addRemoveLinks: true,
      }
    }
  },
  created() {
    this.dropzoneOptions.accept = this.acceptUploadedFile;
    this.dropzoneOptions.removedfile = this.removeUploadedFile;
  },
  methods: {
    questionChanged(val, oldVal) {
      console.log(val);
      console.log(oldVal);
    },
    async acceptUploadedFile(file, done) {
      this.images.push([file.upload.uuid, file]);
      done();
    },
    removeUploadedFile(file) {
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
      if (this.question.length == 0 || this.images.length < 2) {
        return null;
      }
      let tournament = {
        "question": this.question,
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
      if (this.images) {
        const tournament = await this.createTournamentRequest();

        if (tournament == null) {
          console.log("Add more pictures or ask a question");
          return;
        }
  
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