<template>
  <div class="mash_result">
    Result of mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id">Back to mash sample</router-link>
    <hr />
    <div class="results">
      {{ question }} <br />
      <div class="row" v-for="(choice, rank) in choices" :key="rank">
        <div class="col-md-4 offset-md-4">
          {{ rank + 1 }}: {{ choice["elo"] }} <br />
          <b-button v-if="choice['type'] == 1" :pressed="true" variant="outline-primary">
            <b-img thumbnail :src="choice['bytestream']"></b-img>
          </b-button>
          <b-button v-else :pressed="true" variant="outline-primary">{{ choice['bytestream'] }}</b-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

export default {
  name: 'MashResult',
  data() {
    return {
      question: "",
      choices: [],
    }
  },
  async created() {
    try {
      const res = await axios.get("http://localhost:8081/result", {params: {id: this.$route.params.id}})
      console.log(res.data);
      this.question = res.data["question"];
      this.choices = res.data["result"];
    } catch(e) {
      console.error(e);
    }
  }
}
</script>
