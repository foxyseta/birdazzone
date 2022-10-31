<template>
    <Disclosure as="nav" class="bg-cyan-200">
      <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
        <div class="relative flex h-16 items-center justify-between">
          <div class="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
            <div class="flex flex-shrink-0 items-center">
              <img class="block h-8 w-auto lg:hidden" src="/logoFede.jpeg" alt="Your Company" />
              <img class="hidden h-8 w-auto lg:block" src="/logoFede.jpeg" alt="Your Company" />
            </div>
          </div>
        </div>
      </div>
    </Disclosure>
  </template>
  
  <script setup lang=ts>
  import { Disclosure, DisclosureButton, DisclosurePanel, Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
  import { Bars3Icon, BellIcon, XMarkIcon } from '@heroicons/vue/24/outline'
  </script>
