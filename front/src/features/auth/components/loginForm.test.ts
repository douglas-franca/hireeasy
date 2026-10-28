import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import LoginForm from '@/features/auth/components/LoginForm.vue'

describe('LoginForm', () => {
  it('emits validated credentials when submitted', async () => {
    const wrapper = mount(LoginForm, {
      props: { errorMessage: null, isSubmitting: false },
    })
    const inputs = wrapper.findAll('input')

    await inputs[0]?.setValue('ana@example.com')
    await inputs[1]?.setValue('password123')
    await wrapper.find('form').trigger('submit')

    expect(wrapper.emitted('submit')).toEqual([[{ email: 'ana@example.com', password: 'password123' }]])
  })
})
