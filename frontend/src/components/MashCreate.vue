<template>
  <div class="mash-create">
    <b-row align-h="start">
      <b-col cols="2" align-h="left">
        <router-link icon class="h3" to="/"><b-icon icon="house-door-fill"></b-icon></router-link>
      </b-col>
      <b-col cols="8">
        <b-button variant="outline-primary" @click="startTournament">Let's start!</b-button>
      </b-col>
    </b-row>
    <hr />
    <div class="mash-create-choices">
      <b-container fluid>
        <b-row class="my-1">
          <b-col sm="3" offset="2">
            <label>Question: </label>
          </b-col>
          <b-col sm="5">
            <b-form-input v-model="question" type="text" placeholder="What's best?"></b-form-input>
          </b-col>
        </b-row>
        <br/>
        <label>Possible choices</label>
        <b-tabs>
          <b-tab title="Images" active>
            <vue-dropzone ref="dropzone" id="dropzone" :options="dropzoneOptions"></vue-dropzone>
          </b-tab>
          <b-tab title="Text">
            <b-form-textarea
              id="textarea"
              v-model="textchoices"
              placeholder="Choice 1
Choice 2"
            ></b-form-textarea>
          </b-tab>
        </b-tabs>
      </b-container>
    </div>
  </div>
</template>

<script>
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
      question: "What's the best choice?",
      textchoices: "",
      dropzoneOptions: {
        url: 'https://httpbin.org/post',
        dictDefaultMessage: 'Drop images here to upload',
        resizeWidth: 200,
        resizeHeight: 300,
        maxFilesize: 2,
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
      const tchoices = this.textchoices.split("\n").filter(function(el) {
        return el.length > 0;
      });
      if (this.question.length == 0 || this.images.length + tchoices.length < 2) {
        return null;
      }
      let tournament = {
        "question": this.question,
        "size": this.images.length + tchoices.length,
        "choices": tchoices,
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