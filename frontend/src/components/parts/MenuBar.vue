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
      <a v-else-if="item.command" @click="item.command" v-bind="props.action">
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
import { useRoute, useRouter } from 'vue-router'
import Menubar from 'primevue/menubar'
import { useConfirm } from 'primevue/useconfirm'
import type { MenuItem, MenuItemCommandEvent } from 'primevue/menuitem'
import { fetchApi } from '@/plugins/api'

const { t } = useI18n()
const confirm = useConfirm()
const router = useRouter()

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

const shutdownConfirm = (event: MenuItemCommandEvent) => {
  confirm.require({
    target: event.currentTarget,
    header: t('shutdown-confirm-modal.title'),
    message: t('shutdown-confirm-modal.message'),
    acceptLabel: t('shutdown-confirm-modal.accept'),
    rejectLabel: t('shutdown-confirm-modal.reject'),
    icon: 'pi pi-power-off',
    accept: () => {
      shutdown()
    }
  })
}

const shutdown = async () => {
  try {
    await fetchApi('/shutdown', {
      method: 'POST'
    })
  } catch (e) {
    console.warn(e)
  }

  await router.push({ name: 'goodbye' })
}

const items: MenuItem[] = [
  {
    label: t('menu.search.label'),
    icon: 'pi pi-fw pi-search',
    route: { name: 'search' }
  },
  {
    label: t('menu.files.label'),
    icon: 'pi pi-fw pi-file',
    route: { name: 'files' }
  },
  {
    label: t('menu.tags.label'),
    icon: 'pi pi-fw pi-tag',
    route: { name: 'tags' }
  },
  {
    label: t('menu.info.label'),
    icon: 'pi pi-fw pi-info-circle',
    route: { name: 'info' }
  },
  {
    label: t('menu.settings.label'),
    icon: 'pi pi-fw pi-cog',
    route: { name: 'settings' }
  },
  {
    label: t('menu.quit.label'),
    icon: 'pi pi-fw pi-power-off',
    command: shutdownConfirm
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
