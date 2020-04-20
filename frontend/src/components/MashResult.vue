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
      <transition name="fade-taunt" appear>
        <h3><b>{{ question }}</b></h3>
      </transition>
      <hr />
      <transition-group appear name="res-fade">
        <b-row v-for="(choice, rank) in choices" :key="rank+1" class="mt-4 mx-4" align-v="center">
          <b-col cols='2' offset-md='1' offset-lg='2'>
            <h4><b>{{ rank + 1 }}</b></h4> score: {{ choice["elo"] }}
          </b-col>
          <b-col cols='8' md='6' lg='4'>
            <div class="choice-result" v-if="choice['type'] == 1">
              <b-img class="choice-img" thumbnail :src="choice['bytestream']"></b-img>
            </div>
            <div class="choice-result" variant="outline-primary" v-else>{{ choice['bytestream'] }}</div>
          </b-col>
        </b-row>
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
      const endpoint = process.env.VUE_APP_BACKEND_BASE_ADDRESS + "/api/result";
      const res = await axios.get(endpoint , {params: {id: this.$route.params.id}});
      console.log(res.data);
      this.question = res.data["question"];
      this.choices = res.data["result"];
    } catch(e) {
      if (e.response.status == 404) {
        window.location.replace("#/notfound");
      } else {
        console.error(e);
      }
    }
  }
}
</script>
