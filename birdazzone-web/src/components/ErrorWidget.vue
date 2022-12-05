<script setup lang="ts">
import { ref } from 'vue'
import { Dialog, DialogPanel, DialogTitle } from '@headlessui/vue'

const props = defineProps<{ open:boolean, title: string | undefined, text: string | undefined}>()
const open = ref<boolean>()

</script>

<template>
  <Dialog v-if="open" as="div" class="relative z-10" @close="open = false">
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="flex min-h-full items-end justify-center font-semibold p-4 text-center sm:items-center sm:p-0">
          <DialogPanel class="relative transform overflow-hidden rounded-lg bg-foreground text-left text-white shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
            <div class="px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <div class="sm:flex sm:items-start">
                <div class="mt-3 text-center sm:mt-0 sm:ml-4">
                  <DialogTitle as="h3" class="text-lg font-medium leading-6">{{title}}</DialogTitle>
                  <div class="mt-2">
                    <p class="text-sm text-lgray">{{text}}</p>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-background px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
              <button type="button" class="font-semibold text-white border-b border-lgreen bg-foreground hover:bg-lgreen rounded-lg p-2 text-center text-sm"
                @click="open = false">
                Close
              </button>
            </div>
          </DialogPanel>
        </div>
      </div>
    </Dialog>
</template>
