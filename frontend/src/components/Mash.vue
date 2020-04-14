<template>
  <div class="mash">
    Mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id + '/result'">Result</router-link>
    <hr />
    <transition name="slide-left">
      <b-container ref="questionContainer" v-if="lastOpinion.length > 0" fluid>
        <b-row class="my-1 justify-content-center">{{ lastOpinion }}</b-row>
      </b-container>
    </transition>
    {{ choices["question"] }}
    <div class="row justify-content-md-center">
      <transition appear name="slide-left">
        <div v-if="choiceLoaded" class="col col-md-3 align-middle">
          <b-button  v-if="choices['c1']['type'] == 1" variant="outline-primary" v-on:click="chooseWinner(0)">
            <b-img thumbnail :src="choices['c1']['bytestream']"></b-img>
          </b-button>
          <b-button v-else variant="outline-primary" v-on:click="chooseWinner(0)" >
            {{ choices['c1']['bytestream'] }}
          </b-button>
        </div>
      </transition>
      <div class="col-md-auto">VS</div>
      <transition appear name="slide-right">
        <div v-if="choiceLoaded" class="col col-md-3 align-middle">
          <b-button  v-if="choices['c2']['type'] == 1" variant="outline-primary" v-on:click="chooseWinner(1)">
            <b-img thumbnail :src="choices['c2']['bytestream']"></b-img>
          </b-button>
          <b-button v-else variant="outline-primary" v-on:click="chooseWinner(1)" >
            {{ choices['c2']['bytestream'] }}
          </b-button>
        </div>
      </transition>
    </div>
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
        const res = await axios.post("http://localhost:8081/choice", this.choices, config)
        this.displayNewChoice(res.data);
      } catch(e) {
        console.error(e);
      }
    }
  },
  async created() {
    try {
      const res = await axios.get("http://localhost:8081/choice", {params: {id: this.$route.params.id}})
      console.log(res.data);
      this.choices = res.data;
      this.choiceLoaded = true;
    } catch(e) {
      console.error(e);
    }
  }
}
</script>
