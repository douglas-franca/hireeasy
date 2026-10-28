<script setup lang="ts">
import { reactive, shallowRef } from 'vue'

import { type RegisterFormValues, registerSchema } from '@/features/auth/model/authValidation'
import { Button } from '@/shared/ui/button'
import { Field, FieldDescription, FieldGroup, FieldLabel } from '@/shared/ui/field'
import { Input } from '@/shared/ui/input'
import { Spinner } from '@/shared/ui/spinner'

const props = defineProps<{ errorMessage: string | null; isSubmitting: boolean }>()
const emit = defineEmits<{ submit: [values: RegisterFormValues] }>()

const values = reactive<RegisterFormValues>({ name: '', email: '', password: '' })
const fieldError = shallowRef<string | null>(null)

function submit() {
  const result = registerSchema.safeParse(values)

  if (!result.success) {
    fieldError.value = result.error.issues[0]?.message ?? 'Check the form fields.'
    return
  }

  fieldError.value = null
  emit('submit', result.data)
}
</script>

<template>
  <form class="flex flex-col gap-6" @submit.prevent="submit">
    <FieldGroup>
      <Field :data-invalid="Boolean(fieldError)">
        <FieldLabel for="register-name">Name</FieldLabel>
        <Input
          id="register-name"
          v-model="values.name"
          autocomplete="name"
          :aria-invalid="Boolean(fieldError)"
        />
      </Field>

      <Field :data-invalid="Boolean(fieldError)">
        <FieldLabel for="register-email">Email</FieldLabel>
        <Input
          id="register-email"
          v-model="values.email"
          type="email"
          autocomplete="email"
          :aria-invalid="Boolean(fieldError)"
        />
      </Field>

      <Field :data-invalid="Boolean(fieldError)">
        <FieldLabel for="register-password">Password</FieldLabel>
        <Input
          id="register-password"
          v-model="values.password"
          type="password"
          autocomplete="new-password"
          :aria-invalid="Boolean(fieldError)"
        />
        <FieldDescription v-if="fieldError || props.errorMessage">{{
          fieldError ?? props.errorMessage
        }}</FieldDescription>
      </Field>
    </FieldGroup>

    <Button type="submit" :disabled="props.isSubmitting">
      <Spinner v-if="props.isSubmitting" data-icon="inline-start" />
      Create account
    </Button>
  </form>
</template>
