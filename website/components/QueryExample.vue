<template>
  <div class="QueryExample">
    <slot/>
    <div class="language-text vp-adaptive-theme">
      <button title="Copy Code" class="copy"></button>
      <span class="lang">query</span>
      <!--  @formatter:off -->
      <pre class="shiki shiki-themes github-light github-dark vp-code"><code><span class="line"><span>{{ written }}</span></span></code></pre>
      <!--  @formatter:on -->
    </div>

    <div v-if="demo" class="DemoLink">
      <a :href="url" rel="noopener noreferrer" target="_blank">Open query in demo</a>
    </div>

  </div>
</template>

<style lang="scss">
.QueryExample {
  display: flex;
  flex-direction: column-reverse;

  .language-text.vp-adaptive-theme {
    margin: 0 !important;
  }

  .DemoLink {
    font-size: .9rem;
    margin-top: .4rem;
    text-align: right;
  }
}

html {
  &.dark {
    .QueryExample {
      border-color: white;
    }
  }
}
</style>

<script setup lang="ts">
import {computed} from "vue";

const props = withDefaults(defineProps<{
  query?: string
  tags?: string[]
  demo?: boolean
}>(), {
  demo: true,
})

const url = computed(() => {
  const q = new URLSearchParams({});

  if (props.query) {
    q.append('q', props.query)
  }

  if (props.tags) {
    props.tags.forEach(tag => {
      q.append('tag', tag)
    })
  }

  return `https://demo.ablegram.app/app/search?${q.toString()}`
})

const written = computed(() => {
  const parts = [];

  if (props.query) {
    parts.push(props.query)
  }

  if (props.tags) {
    props.tags.forEach(tag => {
      // extract the boolean prefix, if present
      const match = tag.match(/^(?<prefix>[-+])?(?<tag>.*)$/)
      const part = (match?.groups?.prefix ?? '') + 'tags:"' + (match?.groups?.tag ?? '') + '"'
      if (part.trim() !== '') {
        parts.push(part)
      }
    })
  }

  return parts.join(' ')
})
</script>
