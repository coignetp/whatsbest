<template>
  <div class="mash">
    Mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id + '/result'">Result</router-link>
    <hr />
    {{ choices["question"] }}
    <div class="row justify-content-md-center">
      <div class="col col-md-3">
        <b-img thumbnail v-on:click="chooseWinner(0)" :src="choices['c1']['bytestream']"></b-img>
      </div>
      <div class="col-md-auto">VS</div>
      <div class="col col-md-3">
        <b-img thumbnail v-on:click="chooseWinner(1)" :src="choices['c2']['bytestream']"></b-img>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'


export default {
  name: 'Mash',
  data() {
    return {
      choices: {
        "c1": {
          "type": 1,
          "bytestream": ""
        },
        "c2": {
          "type": 1,
          "bytestream": ""
        }
      }
    }
  },
  methods: {
    displayNewChoice(data) {
      console.log(data);
      this.choices = data;
    },

    async chooseWinner(i) {
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
      this.displayNewChoice(res.data);
    } catch(e) {
      console.error(e);
    }
  }
}
</script>
