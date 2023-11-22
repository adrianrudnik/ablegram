<script lang="ts">
import {h} from "vue";

export default {
  setup(props, {slots}) {
    if (!slots || !slots.default) {
      return () => null;
    }

    const content = slots.default();

    return () =>
        h("script", {
          innerHTML: content.map(el => el.children).join(''),
          type: "application/ld+json",
        });
  },
};
</script>
