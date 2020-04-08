<template>
  <div class="big_button">
    <button v-on:click="counter += 1">Add 1</button>
    <p>Le bouton ci-dessus a été cliqué {{ counter }} fois.</p>
  </div>
</template>

<script>
export default {
  name: 'Home',
  props: {
    counter: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
