<template>
  <div class="mash">
    Mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id + '/result'">Result</router-link>
    <hr />
    <div class="choice left">
      <b-img left id="ChoiceLeft" v-on:click="chooseWinner(0)" src=""></b-img>
    </div>
    <div class="choice right">
      <b-img right id="ChoiceRight" v-on:click="chooseWinner(1)" src=""></b-img>
    </div>
  </div>
</template>

<script>
import axios from 'axios';


export default {
  name: 'Mash',
  methods: {
    chooseWinner(i) {
      this.choices["winner"] = i;

      const config = {
        headers: {
          'Content-Type': 'application/json',
        }
      };

      // 
      console.log("Winner " + i + " will be sent");

      this.choices["c1"]["bytestream"] = "";
      this.choices["c1"]["idtournament"] = parseInt(this.$route.params.id, 16);
      this.choices["c2"]["bytestream"] = "";
      this.choices["c2"]["idtournament"] = parseInt(this.$route.params.id, 16);

      axios.post("http://localhost:8081/choice", this.choices, config).then(response=>{
        console.log(response.data);
        this.choices = response.data;
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
  },
  beforeMount() {
    axios.get("http://localhost:8081/choice", {params: {id: this.$route.params.id}}).then(response=>{
      console.log(response.data);
      this.choices = response.data;
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
