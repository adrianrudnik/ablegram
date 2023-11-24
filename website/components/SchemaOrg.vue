<script lang="ts">
import {h} from "vue";

export default {
  setup(props, {slots}) {
    if (!slots?.default) {
      return () => null;
    }

    return () =>
        h("script", {
          innerHTML: slots.default?.().map(el => el.children).join(''),
          type: "application/ld+json",
        });
  },
};
</script>
