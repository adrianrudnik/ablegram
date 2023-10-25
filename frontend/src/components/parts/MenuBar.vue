<template>
  <Menubar :model="items" :pt="pt">
    <template #item="{ label, item, props, root, hasSubmenu }">
      <RouterLink v-if="item.route" v-slot="{ navigate, href }" :to="item.route" custom>
        <a
          :href="href"
          v-bind="props.action"
          @click="navigate"
          :class="{ 'active-link border-round-md': activeRoute === item.route?.name }"
          class="border-transparent"
        >
          <span v-bind="props.icon" />
          <span v-bind="props.label">{{ label }}</span>
        </a>
      </RouterLink>
      <a v-else :href="item.url" :target="item.target" v-bind="props.action">
        <span v-bind="props.icon" />
        <span v-bind="props.label">{{ label }}</span>
        <span
          :class="[hasSubmenu && (root ? 'pi pi-fw pi-angle-down' : 'pi pi-fw pi-angle-right')]"
          v-bind="props.submenuicon"
        />
      </a>
    </template>
  </Menubar>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router'
import Menubar from 'primevue/menubar'
import type { MenuItem } from 'primevue/menuitem'

const { t } = useI18n()

const activeRoute = computed(() => useRoute().name)

const pt = {
  root: {
    class: 'bg-white border-none pl-0'
  },
  menuitem: {
    class: 'pr-1'
  },
  content: {
    class: 'border-1'
  }
}

const items: MenuItem[] = [
  {
    label: t('menu.search.label'),
    icon: 'pi pi-fw pi-search',
    route: { name: 'search' }
  },
  {
    label: t('menu.status.label'),
    icon: 'pi pi-fw pi-bolt',
    route: { name: 'status' }
  },
  {
    label: t('menu.about.label'),
    icon: 'pi pi-fw pi-info-circle',
    route: { name: 'about' }
  },
  {
    label: t('menu.quit.label'),
    icon: 'pi pi-fw pi-power-off'
  }
]
</script>

<style lang="scss">
.p-menuitem-link.active-link {
  background-color: var(--surface-900);

  .p-menuitem-text,
  .p-menuitem-icon {
    color: white;
  }

  font-weight: bold;
}
</style>
