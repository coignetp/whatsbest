<template>
  <div class="mash_result">
    Result of mashing n°{{ $route.params.id }} <br />
    <router-link :to="'/mash/' + $route.params.id">Back to mash sample</router-link>
  </div>
</template>

<script>
export default {
  name: 'MashResult',
}
</script>
