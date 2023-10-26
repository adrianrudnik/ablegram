<template>
  <li class="MetaStepItem relative py-3 pl-4 pr-3 flex flex-column justify-content-center">
    <div
      class="absolute left-0 top-0 z-1"
      style="
        border-left: 20px solid var(--surface-900);
        border-top: 45px solid transparent;
        border-bottom: 45px solid transparent;
        width: 0;
        height: 0;
      "
    ></div>
    <div class="text-xl font-medium text-white mb-1">
      <slot name="number" />
    </div>

    <span class="text-white white-space-nowrap">
      <slot name="label" />
    </span>

    <div class="Separator"></div>
  </li>
</template>

<style lang="scss"></style>
