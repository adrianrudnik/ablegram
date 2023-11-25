import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import uniqolor from 'uniqolor'

export type UserRoles = 'admin' | 'user'

export interface User {
  id: string
  display_name: string
  role: UserRoles
  ip?: string
}

export const useUserStore = defineStore('users', setupStore<User>())

export const colorizeUser = (
  user: User
): {
  color: string
  isLight: boolean
} => uniqolor(user.id, { format: 'rgb', lightness: 50 })
