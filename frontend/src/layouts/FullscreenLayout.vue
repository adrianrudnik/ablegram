<template>
  <div class="layout">
    <div class="fullscreen">
      <div class="w-25rem">
        <h1 class="text-3xl flex flex-row gap-2 font-medium">
          <img src="@/assets/media/logo.svg" class="h-2rem" alt="Ablegram logo" />
          Ablegram
        </h1>

        <RouterView />
      </div>
    </div>
  </div>
</template>

<style lang="scss">
.layout {
  padding: 1rem;
  height: 100%;
  background-color: white;

  .fullscreen {
    min-height: 100vh;
    padding: 0;
    display: flex;
    flex-wrap: nowrap;
    justify-content: center;
    align-items: center;

    a {
      color: black;

      &:hover {
        color: black;
      }
    }
  }

  &.h-center {
    .fullscreen {
      justify-content: center;
    }
  }
}
</style>
