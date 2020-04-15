<template>
  <b-row align-h="start" align-v="center" class="mt-5">
    <b-col cols="2" align-h="left">
      <router-link icon class="h3" to="/"><b-icon icon="house-door-fill"></b-icon></router-link>
    </b-col>
    <b-col cols="8">
      <h1>This page does not exist.</h1>
    </b-col>
  </b-row>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

export default {
  name: 'NotFound'
}
</script>