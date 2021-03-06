<template>
  <div class="mash-create">
    <b-row align-h="start">
      <b-col cols="2" align-h="left">
        <router-link icon class="h3" to="/"><b-icon icon="house-door-fill"></b-icon></router-link>
      </b-col>
      <b-col cols="8">
        <b-button variant="outline-primary" @click="startTournament">Let's start!</b-button>
      </b-col>
      <notifications group="missing_choices" />
      <notifications group="missing_question" />
    </b-row>
    <hr />
    <div class="mash-create-choices">
      <b-container fluid>
        <b-row class="my-4" align-v="center">
          <b-col sm="3" offset-sm="2">
            <label>Question: </label>
          </b-col>
          <b-col sm="5">
            <b-form-input v-model="question" type="text" placeholder="What's best?"></b-form-input>
          </b-col>
        </b-row>
        <br/>
        <label>Possible choices:</label>
        <b-tabs nav-class="ml-1">
          <b-tab title="Images" active>
            <vue-dropzone ref="dropzone" id="dropzone" :options="dropzoneOptions" style="border-radius:5px"></vue-dropzone>
          </b-tab>
          <b-tab title="Text">
            <b-form-textarea
              id="textarea"
              v-model="textchoices"
              :placeholder=textPlaceholder
            ></b-form-textarea>
          </b-tab>
        </b-tabs>
      </b-container>
    </div>
    <b-row class="mt-5" align-h="center"><b-col cols="8">
      <b-button variant="outline-primary" @click="startTournament">Let's start!</b-button>
    </b-col></b-row>
  </div>
</template>

<script>
import vue2Dropzone from "vue2-dropzone";
import 'vue2-dropzone/dist/vue2Dropzone.min.css';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import './../style.css';

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
      textPlaceholder: "Choice 1\nChoice 2",
      images: [],
      question: "What's the best choice?",
      textchoices: "",
      dropzoneOptions: {
        // url to accept the uploaded file
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
    // Link the dropzone with custom methods after the component creation
    this.dropzoneOptions.accept = this.acceptUploadedFile;
    this.dropzoneOptions.removedfile = this.removeUploadedFile;
  },
  methods: {
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
      // The question field must not be empty
      if (this.question.length == 0) {
        this.$notify({
          group: "missing_question",
          type: "error",
          title: "Missing question",
          text: "Complete the 'Question' field before starting the tournament."
        });
        return null;
      }
      // At least 2 choices (images + text)
      if (this.images.length + tchoices.length < 2) {
        this.$notify({
          group: "missing_choices",
          type: "error",
          title: "Not enough choices",
          text: "You need at least 2 different choices to start the tournament. In Text, you can separate them with newlines."
        });
        return null;
      }
      let tournament = {
        "question": this.question,
        "size": this.images.length + tchoices.length,
        "choices": tchoices,
      };

      // Convert images to base64 before sending it to the backend
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
          const endpoint = process.env.VUE_APP_BACKEND_BASE_ADDRESS + "/api/create";
          const res = await axios.post(endpoint, tournament, config);
          if (res.status == 200) {
            // Once the tournament is created, the user is sent to the choices
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
