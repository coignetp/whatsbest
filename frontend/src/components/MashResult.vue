<template>
  <div class="mash_result">
    <b-row align-h="start">
      <b-col cols="2" align-h="left">
        <router-link icon class="h3" to="/"><b-icon icon="house-door-fill"></b-icon></router-link>
      </b-col>
      <b-col cols="8">
        <b-button variant="outline-primary" :to="'/mash/' + $route.params.id">Back to the decisions!</b-button>
      </b-col>
    </b-row>
    <hr />
    <div class="results">
      {{ question }}
      <hr />
      <transition-group appear name="slide-right">
        <div class="row" v-for="(choice, rank) in choices" :key="`key-${rank}`">
          <div class="col-md-4 offset-md-4">
            {{ rank + 1 }}: {{ choice["elo"] }} <br />
            <b-button v-if="choice['type'] == 1" :pressed="true" variant="outline-primary">
              <b-img thumbnail :src="choice['bytestream']"></b-img>
            </b-button>
            <b-button v-else :pressed="true" variant="outline-primary">{{ choice['bytestream'] }}</b-button>
          </div>
        </div>
      </transition-group>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import './../style.css';

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
