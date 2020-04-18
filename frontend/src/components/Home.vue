<template>
  <div class="big_button">
    <b-container class="pt-5 justify-content-center">
      <b-row align-h="center" class="my-5 mx-3"><h4>You wanna rank some stuff? Do it here!</h4></b-row>
      <b-row align-h="center">
        <b-button variant="outline-primary" to="/mash/create"><h4>Start a tournament</h4></b-button>
      </b-row>
    </b-container>
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
</script>