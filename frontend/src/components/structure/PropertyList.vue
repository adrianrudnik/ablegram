<template>
  <div class="PropertyList">
    <ul class="list-none p-0 m-0">
      <slot />
    </ul>
  </div>
</template>

<style lang="scss">
.PropertyList {
  .PropertyListItem {
    padding-bottom: 0.5rem !important;

    &:not(:first-child) {
      border-top: 1px solid var(--gray-400);
      padding-top: 0.5rem !important;
    }
  }
}
</style>
