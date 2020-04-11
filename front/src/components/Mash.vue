<template>
  <div class="mash">
    Mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id + '/result'">Result</router-link>
    <hr />
    <div class="choice left">
      <img id="ChoiceLeft" src="" />
    </div>
    <div class="choice right">
      <img id="ChoiceRight" src="" />
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Mash',
  // methods: {
  //   getChoices() {
  //     axios.get("http://localhost:8081/choice", this.$route.params.id).then(response=>{
  //       console.log(response);
  //     }).catch(err=>{
  //       console.error(err);
  //     });
  //   }
  // },
  beforeMount() {
    axios.get("http://localhost:8081/choice", {params: {id: this.$route.params.id}}).then(response=>{
      console.log(response.data);
      if (response.data["c1"]["type"] === 1) {
        document.getElementById("ChoiceLeft").src = "data:image/png;base64," + response.data["c1"]["bytestream"];
      } else {
        document.getElementById("ChoiceLeft").src = "https://giphy.com/gifs/90s-cartoon-cartoons-Ala8Pjo4RN9kY"
      }
      if (response.data["c2"]["type"] === 1) {
        document.getElementById("ChoiceRight").src = "data:image/png;base64," + response.data["c2"]["bytestream"];
      } else {
        document.getElementById("ChoiceLeft").src = "https://giphy.com/gifs/chillin-n9j743yPQ5pg4"
      }
  }).catch(err=>{
      console.error(err);
    });
  }
}
</script>
