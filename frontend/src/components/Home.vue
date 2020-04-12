<template>
  <div class="big_button">
    Welcome <br />
    
    <router-link to="/mash/create" tag="button">Start a tournament</router-link>
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
