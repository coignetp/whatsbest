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
    <transition name="fade-taunt">
      <b-container ref="questionContainer" v-if="lastOpinion.length > 0" fluid>
        <b-row align-v="center" align-h="center"><h5>{{ lastOpinion }}</h5></b-row>
      </b-container>
    </transition>
    <hr v-if="lastOpinion.length > 0" />
    <transition name="fade-taunt" appear>
      <h3><b>{{ choices["question"] }}</b></h3>
    </transition>
    <hr/>
    <!-- Display the 2 possible choices according to their type. 1 is image, 0 is other. -->
    <b-row class="mt-5" align-v="center" align-h="center">
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
      <transition name="fade-taunt">
        <b-col cols="auto">
          <h1 v-if="choiceLoaded"><b>VS</b></h1>
          <div v-if="!choiceLoaded" class="spinner-border text-primary" role="status">
            <span class="sr-only">Loading...</span>
          </div>
        </b-col>
      </transition>
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
      // Little taunt after the user made a choice
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
        "If you say so",
        "Are you sure?",
        "You chose poorly",
        "You have chosen... wisely",
        "This one was easy",
        "The next one is easier",
        "You did your best",
      ],
    }
  },
  methods: {
    displayNewChoice(data) {
      console.log(data);
      this.choices = data;
      this.choiceLoaded = true; 
    },

    async chooseWinner(i) {
      // Send the winner to the backend
      this.lastOpinion = ""
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
        this.lastOpinion = this.opinions[Math.floor(Math.random() * this.opinions.length)];
        const endpoint = process.env.VUE_APP_BACKEND_BASE_ADDRESS + "/api/choice";
        // Once the winner is sent, the frontend receive a new choice to display
        const res = await axios.post(endpoint, this.choices, config);
        this.displayNewChoice(res.data);
      } catch(e) {
        console.error(e);
      }
    }
  },
  async created() {
    try {
      const endpoint = process.env.VUE_APP_BACKEND_BASE_ADDRESS + "/api/choice";
      const res = await axios.get(endpoint, {params: {id: this.$route.params.id}})
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
