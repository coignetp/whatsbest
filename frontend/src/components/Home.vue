<template>
  <div class="big_button">
    Welcome <br />
    <b-button variant="outline-primary" to="/mash/create">Start a tournament</b-button>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

export default {
  name: 'Home',
  props: {
    counter: String
  }
}

