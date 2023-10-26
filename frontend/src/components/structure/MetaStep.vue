<template>
  <ul
    class="MetaStep list-none p-0 m-0 surface-900 flex overflow-y-hidden overflow-x-auto border-round"
  >
    <slot />
  </ul>
</template>

<style lang="scss">
.MetaStep {
  .MetaStepItem:first-child {
    .Separator {
      display: none;
    }
  }
  .MetaStepItem:not(:first-child) {
    padding-left: 3rem !important;

    .Separator {
      position: absolute;
      top: 0;
      left: 1px;
      border-left: 20px solid var(--surface-300);
      border-top: 45px solid transparent;
      border-bottom: 45px solid transparent;
      width: 0;
      height: 0;
    }
  }
}
</style>
