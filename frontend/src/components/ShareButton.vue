<template>
  <div>
    <b-button variant="outline-primary" v-b-modal.modal-share>Share it</b-button>

    <b-modal ok-only id="modal-share" title="Send this link to your friends!">
      <b-col align-h="center" class="my-2"><i>{{ modalLink }}</i></b-col>
    </b-modal>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

export default {
  name: 'ShareButton',
  props: {
    modalLink: String
  },
  data() {
    return {
      showModal: false,
    }
  }
}
</script>