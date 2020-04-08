<template>
  <div class="mash">
    Mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id + '/result'">Result</router-link>
  </div>
</template>

<script>
export default {
  name: 'Mash',
}
</script>
