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

export default {
  name: 'MashCreate',
  components: {
    PictureInput
  },
  methods: {
    onChange () {
      console.log('New picture selected!')
      if (this.$refs.pictureInput.image) {
        console.log('Picture loaded.')
        this.image = [
          this.$refs.pictureInput.file,
          this.$refs.pictureInput.file,
        ];
      } else {
        console.log('FileReader API not supported: use the <form>, Luke!')
      }
    },

    startTournament() {
      console.log("In start Tournament");
      if (this.image) {
        const formData = new FormData();
        formData.append("question", "What's best?");
        formData.append("size", 2);
        formData.append("image0", this.image[0]);
        formData.append("image1", this.image[1]);
        const config = {
          headers: {
            'content-type': 'multipart/form-data'
          }
        };

        axios.post("http://localhost:8081/create", formData, config).then(response=>{
          if (response.data.success){
            this.image = '';
            console.log("Image uploaded successfully ✨");
          }
        }).catch(err=>{
          console.error(err);
        });
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