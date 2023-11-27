import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import uniqolor from 'uniqolor'

export const AdminRole = 'admin'
export type AdminRole = typeof AdminRole
export const GuestRole = 'guest'
export type GuestRole = typeof GuestRole

export type UserRoles = AdminRole | GuestRole

export interface UserClient {
  id: string
  ip?: string

  user_id: string
  user_display_name: string
  user_role: UserRoles
}

export const useUserClientStore = defineStore('clients', setupStore<UserClient>())

export const colorizeUserClient = (
  user: UserClient
): {
  color: string
  isLight: boolean
} => uniqolor(user.id, { format: 'rgb', lightness: 50 })
