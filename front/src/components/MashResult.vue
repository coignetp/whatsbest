<template>
  <div class="mash_result">
    Result of mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id">Back to mash sample</router-link>
    <hr />
    <div class="results">
      {{ question }} <br />
      <div class="choice_res" v-for="(choice, rank) in choices" :key="rank">
        {{ rank + 1 }}: {{ choice["elo"] }} <br />
        <b-img :src="choice['bytestream']"></b-img>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'MashResult',
  data() {
    return {
      question: "",
      choices: [],
    }
  },
  beforeCreate() {
    axios.get("http://localhost:8081/result", {params: {id: this.$route.params.id}}).then(response=>{
      console.log(response.data);
      this.question = response.data["question"];
      this.choices = response.data["result"];
      console.log(this.choices);
    }).catch(err=>{
      console.error(err);
    });
  }
}
</script>
