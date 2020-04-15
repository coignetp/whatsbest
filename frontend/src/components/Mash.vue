<template>
  <div class="mash">
    <b-row align-h="start">
      <b-col cols="2" align-h="left">
        <router-link icon class="h3" to="/"><b-icon icon="house-door-fill"></b-icon></router-link>
      </b-col>
      <b-col cols="8">
        <b-button variant="outline-primary" :to="'/mash/' + $route.params.id + '/result'">Current ranking</b-button>
      </b-col>
    </b-row>
    <hr />
    <transition name="slide-left">
      <b-container ref="questionContainer" v-if="lastOpinion.length > 0" fluid>
        <b-row align-v="center" align-h="center">{{ lastOpinion }}</b-row>
      </b-container>
    </transition>
    <hr v-if="lastOpinion.length > 0" />
    <h3><b>{{ choices["question"] }}</b></h3>
    <hr/>
    <b-row align-v="center" align-h="center">
      <b-col col sm="5" lg="4" xl="3">
        <transition appear name="slide-left">
          <div v-if="choiceLoaded" class="col align-middle">
            <b-button  v-if="choices['c1']['type'] == 1" variant="outline-primary" v-on:click="chooseWinner(0)">
              <b-img thumbnail :src="choices['c1']['bytestream']"></b-img>
            </b-button>
            <b-button v-else variant="outline-primary" v-on:click="chooseWinner(0)" >
              {{ choices['c1']['bytestream'] }}
            </b-button>
          </div>
        </transition>
      </b-col>
      <b-col cols="auto"><h1><b>VS</b></h1></b-col>
      <b-col col sm="5" lg="4" xl="3">
        <transition appear name="slide-right">
          <div v-if="choiceLoaded" class="col align-middle">
            <b-button  v-if="choices['c2']['type'] == 1" variant="outline-primary" v-on:click="chooseWinner(1)">
              <b-img thumbnail :src="choices['c2']['bytestream']"></b-img>
            </b-button>
            <b-button v-else variant="outline-primary" v-on:click="chooseWinner(1)" >
              {{ choices['c2']['bytestream'] }}
            </b-button>
          </div>
        </transition>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import './../style.css';


export default {
  name: 'Mash',
  data() {
    return {
      lastOpinion: "",
      choiceLoaded: false,
      choices: {
        "c1": {
          "type": 1,
          "bytestream": ""
        },
        "c2": {
          "type": 1,
          "bytestream": ""
        }
      },
      opinions: [
        "Wise choice",
        "Meh..",
        "Did you read the question?",
        "Why not",
      ],
    }
  },
  methods: {
    displayNewChoice(data) {
      console.log(data);
      this.choices = data;
      this.lastOpinion = this.opinions[Math.floor(Math.random() * this.opinions.length)];
      this.choiceLoaded = true; 
    },

    async chooseWinner(i) {
      this.choiceLoaded = false;
      this.choices["winner"] = i;

      const config = {
        headers: {
          'Content-Type': 'application/json',
        }
      };
      console.log("Winner " + i + " will be sent");

      this.choices["c1"]["bytestream"] = "";
      this.choices["c1"]["idtournament"] = parseInt(this.$route.params.id, 16);
      this.choices["c2"]["bytestream"] = "";
      this.choices["c2"]["idtournament"] = parseInt(this.$route.params.id, 16);

      try {
        const res = await axios.post("/api/choice", this.choices, config)
        this.displayNewChoice(res.data);
      } catch(e) {
        console.error(e);
      }
    }
  },
  async created() {
    try {
      const res = await axios.get("/api/choice", {params: {id: this.$route.params.id}})
      console.log(res.data);
      this.choices = res.data;
      this.choiceLoaded = true;
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
