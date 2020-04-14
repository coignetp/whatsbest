<template>
  <div class="app">
    <div class="nav">
      <router-link icon class="h3 mb-2" to="/"><b-icon icon="house-door-fill"></b-icon></router-link>
    </div>
    <div id="core-view">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

export default {
  name: 'App',
  components: {
  }
}
</script>

<style>
#core-view {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
