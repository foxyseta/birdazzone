<script lang="ts">
import { checkCompatEnabled } from '@vue/compiler-core';
import { ref } from 'vue';
import { defineComponent } from 'vue';
import DatePicker from 'vue-datepicker-next';
import 'vue-datepicker-next/index.css';
import ErrorWidget from './ErrorWidget.vue';

export default defineComponent({
  components: { DatePicker },

  data() {
    return {
      date: null /** to save date values */,
      choosenDate: false /** to verify if date has been chosen */,
      openD: false /** to close date popup */,

      openClose: false /** to open/close filters card */,
    };
  },

  methods: {
    /** GENERAL */
    openCloseFunction() {
      /** to open/close filters card */
      this.openClose = !this.openClose;
      this.date = null;
      this.choosenDate = false;
    },
    sendData() {
      /** confirm button */
      this.openCloseFunction();
    },

    /** DATE */
    closeD() {
      /** to close date popup */
      this.openD = false;
    },
    disabledDate(date: Date) {
      /** to disabilitate days after today and before last week in date popup */
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      return date > today || date < new Date(today.getTime() - 6 * 24 * 3600 * 1000);
    },
    selectDate() {
      /** to confirm selected date */
      if (this.date != null) {
        this.choosenDate = true;
      } // ELSE -> didnt enter date -> by default: today
    },
    modifyDate() {
      this.choosenDate = false;
    },
  },
});
</script>

<template>
  <div class="flex justify-start">
    <button
      class="font-semibold text-white border border-lgreen bg-foreground hover:bg-lgreen rounded-lg px-4 py-2 text-center text-sm inline-flex items-center"
      type="button"
      @click="openCloseFunction()"
    >
      <svg
        class="mr-2 w-4 h-4"
        alt="arrow down"
        aria-hidden="true"
        fill="none"
        stroke="white"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
      </svg>
      Date filter
    </button>
  </div>

  <div
    v-if="openClose"
    class="z-10 bg-foreground border border-background p-4 font-semibold text-md rounded-lg m-2 mr-0 place-self-start"
  >
    <div class="flex flex-row items-stretch justify-center" aria-labelledby="dropdownDividerButton">
      <label for="Date" class="justify-self-start self-center text-white text-sm w-9">date</label>
      <div class="flex justify-self-start" id="Date">
        <date-picker
          class="m-2 mr-0"
          type="date"
          v-model:value="date"
          v-model:open="openD"
          value-type="format"
          format="YYYY-MM-DD"
          placeholder="select date"
          :clearable="true"
          :disabled-date="disabledDate"
          :disabled="choosenDate"
        >
          <template #footer>
            <button class="mx-btn mx-btn-text" @click="closeD()">close</button>
          </template>
        </date-picker>
      </div>
    </div>
    <div v-if="choosenDate" class="flex justify-center">
      <button
        class="font-semibold text-white hover:text-lgreen text-center text-sm"
        type="button"
        @click="modifyDate()"
      >
        modify date
      </button>
    </div>
    <div v-else class="flex justify-center">
      <button
        class="font-semibold text-white hover:text-lgreen text-center text-sm"
        type="button"
        @click="selectDate()"
      >
        select date
      </button>
    </div>

    <div v-if="choosenDate" class="flex justify-end">
      <button
        class="font-semibold text-white border-b border-lgreen bg-foreground hover:bg-lgreen rounded-lg p-2 text-center text-sm"
        type="button"
        @click="
          $emit('changeDate', date);
          sendData();
        "
      >
        confirm
      </button>
    </div>
  </div>
</template>
