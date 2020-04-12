<template>
  <div class="mash-create">
    <router-link to="/mash/42/result" tag="button">Let's start</router-link> <br />
    <button @click="startTournament">Test start</button>
    <hr />
    <div class="mash-create-choices">
      <picture-input 
        ref="pictureInput" 
        @change="onChange" 
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
    </div>
  </div>
</template>

<script>
import PictureInput from 'vue-picture-input'
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
    PictureInput
  },
  data() {
    return {
      images: []
    }
  },
  methods: {
    onChange () {
      if (this.$refs.pictureInput.image) {
        console.log('Picture loaded.')
        this.images = [
          this.$refs.pictureInput.file,
          this.$refs.pictureInput.file,
        ];
      } else {
        console.log('FileReader API not supported: use the <form>, Luke!')
      }
    },

    async startTournament() {
      console.log("In start Tournament");
      if (this.images) {
        let tournament = {
          "question": "What's best?",
          "size": this.images.length,
          "type": 1,
          "choices": [],
        }

        let img;
        for (img of this.images) {
          tournament["choices"].push(await toBase64(img));
        }

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