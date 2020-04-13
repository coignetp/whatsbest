<template>
  <div class="app">
    <div class="nav">
      <router-link to="/">Home</router-link>
    </div>
    <div id="core-view">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
// import Home from './components/Home.vue'

export default {
  name: 'App',
  components: {
    // Home
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
