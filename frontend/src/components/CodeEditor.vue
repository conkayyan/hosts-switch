<template>
  <codemirror
      v-model="code"
      placeholder="# e.g.
127.0.0.1 localhost"
      :style="{ height: '600px', width: '630px' }"
      :autofocus="true"
      :indent-with-tab="true"
      :tab-size="2"
      :extensions="extensions"
      @ready="log('ready', $event)"
      @change="log('change', $event)"
      @focus="log('focus', $event)"
      @blur="log('blur', $event)"
  />
</template>

<script>
import { ref } from 'vue'
import { Codemirror } from 'vue-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import { oneDark } from '@codemirror/theme-one-dark'

export default {
  components: {
    Codemirror
  },
  setup() {
    const code = ref(`console.log('Hello, world!')`)
    const extensions = [javascript(), oneDark]

    return {
      code,
      extensions,
      log: console.log
    }
  }
}
</script>