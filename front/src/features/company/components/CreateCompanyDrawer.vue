<script setup lang="ts">
import { reactive, shallowRef } from 'vue'

import {
  type CreateCompanyFormValues,
  createCompanySchema,
} from '@/features/company/model/companyValidation'
import { Button } from '@/shared/ui/button'
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
} from '@/shared/ui/drawer'
import { Field, FieldDescription, FieldGroup, FieldLabel } from '@/shared/ui/field'
import { Input } from '@/shared/ui/input'
import { Spinner } from '@/shared/ui/spinner'

const props = defineProps<{
  errorMessage: string | null
  isOpen: boolean
  isSubmitting: boolean
}>()

const emit = defineEmits<{
  close: []
  submit: [values: CreateCompanyFormValues]
}>()

const values = reactive<CreateCompanyFormValues>({ name: '' })
const fieldError = shallowRef<string | null>(null)

function submit() {
  const result = createCompanySchema.safeParse(values)

  if (!result.success) {
    fieldError.value = result.error.issues[0]?.message ?? 'Check the company name.'
    return
  }

  fieldError.value = null
  emit('submit', result.data)
}

function updateOpen(isOpen: boolean) {
  if (!isOpen) emit('close')
}
</script>

<template>
  <Drawer :open="props.isOpen" @update:open="updateOpen">
    <DrawerContent>
      <div class="mx-auto w-full max-w-md">
        <DrawerHeader>
          <DrawerTitle>Create your company</DrawerTitle>
          <DrawerDescription>Set up the workspace where you will manage your hiring.</DrawerDescription>
        </DrawerHeader>

        <form class="px-4" @submit.prevent="submit">
          <FieldGroup>
            <Field :data-invalid="Boolean(fieldError)">
              <FieldLabel for="company-name">Company name</FieldLabel>
              <Input
                id="company-name"
                v-model="values.name"
                autocomplete="organization"
                :aria-invalid="Boolean(fieldError)"
              />
              <FieldDescription v-if="fieldError || props.errorMessage">{{
                fieldError ?? props.errorMessage
              }}</FieldDescription>
            </Field>
          </FieldGroup>

          <DrawerFooter class="px-0">
            <Button type="submit" :disabled="props.isSubmitting">
              <Spinner v-if="props.isSubmitting" data-icon="inline-start" />
              Create company
            </Button>
          </DrawerFooter>
        </form>
      </div>
    </DrawerContent>
  </Drawer>
</template>
