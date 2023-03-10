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
        date: null /** to save the date */,
        openD: false /** to close date popup */,

        sTime: null as string | null /** to same start time value */,
        openST: false /** to close start time popup */,
        eTime: null as string | null /** to same end time value */,
        openET: false /** to close end time popup */,

        sDateTime: null as string | null /** to save start date-time*/,
        eDateTime: null as string | null /** to save end date-time */,

        choosenDate: false /** to verify if date has been chosen */,
        choosenDateTimes: false /** to verify if times and date have been chosen */,
        openClose: false /** to open/close filters card */
      };
    },

    methods: {
      /** GENERAL */
      openCloseFunction() {
        /** to open/close filters card */
        this.openClose = !this.openClose;
        if (this.openClose) {
          this.date = null;
          this.sTime = null;
          this.eTime = null;
          this.sDateTime = null;
          this.eDateTime = null;
          this.choosenDate = false;
          this.choosenDateTimes = false;
        }
      },

      /** DATE */
      closeD() {
        /** to close dates popup */
        this.openD = false;
      },
      disabledDates(date: Date) {
        /** to disabilitate days after today and before last week in dates popup */
        const today = new Date();
        today.setHours(0, 0, 0, 0);

        return date > today || date < new Date(today.getTime() - 6 * 24 * 3600 * 1000);
      },
      selectDate() {
        /** to confirm date was select */
        if (this.date != null) this.choosenDate = true;
        else this.choosenDate = false;
      },
      modifyDate() {
        this.choosenDate = false;
        this.choosenDateTimes = false;
      },

      /** TIMES */
      closeST() {
        /** to close start time popup */
        this.openST = false;
      },
      closeET() {
        /** to close end time popup */
        this.openET = false;
      },
      disabledHours(date: Date) {
        return date.getHours() < 19 || date.getHours() > 20;
      },
      selectTimes() {
        if (this.date !== null && this.sTime !== null && this.eTime !== null) {
          if (this.sTime.substring(0, 2) === this.eTime.substring(0, 2)) {
            // sTime HH == eTime HH (same hours)
            if (this.sTime.substring(3) > this.eTime.substring(3)) {
              // sTime mm > eTime mm -> switch of the two dates
              let box = this.sTime;
              this.sTime = this.eTime;
              this.eTime = box;
            }
          } else {
            // sTime HH != eTime HH (different hours)
            if (this.sTime.substring(0, 2) > this.eTime.substring(0, 2)) {
              // sTime HH > eTime HH -> switch of the two
              let box = this.sTime;
              this.sTime = this.eTime;
              this.eTime = box;
            }
          }

          // the API uses english time zone
          this.sDateTime =
            ((this.date + 'T' + (Number(this.sTime.substring(0, 2)) - 1)) as string) + this.sTime.substring(2) + ':00Z';
          this.eDateTime =
            ((this.date + 'T' + (Number(this.eTime.substring(0, 2)) - 1)) as string) + this.eTime.substring(2) + ':00Z';

          this.choosenDateTimes = true;
        }
        // ELSE -> you must pick a date and both start and end times
      },
      modifyTimes() {
        this.choosenDateTimes = false;
      },

      /** GENERAL */
      sendData() {
        /** confirm button */
        this.openCloseFunction();
      }
    }
  });
</script>

<template>
  <div class="flex justify-start">
    <button
      class="font-semibold text-white text-sm border border-lgreen bg-foreground hover:bg-lgreen rounded-lg px-4 py-2 text-center inline-flex items-center"
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
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 9l-7 7-7-7"
        ></path>
      </svg>
      Date filter
    </button>
  </div>

  <div
    v-if="openClose"
    class="place-self-start z-10 bg-foreground shadow font-semibold text-md rounded-lg m-4 ml-2 p-4"
  >
    <div class="flex flex-row items-start justify-center">
      <label
        for="Date"
        class="justify-self-center self-center text-sm text-white w-7"
        >date</label
      >
      <div
        class="flex justify-self-center"
        id="Date"
      >
        <date-picker
          class="m-2 ml-3"
          type="date"
          v-model:value="date"
          v-model:open="openD"
          value-type="format"
          format="YYYY-MM-DD"
          placeholder="select the date"
          :clearable="true"
          :disabled-date="disabledDates"
          :disabled="choosenDate"
        >
          <template #footer>
            <button
              class="mx-btn mx-btn-text"
              @click="closeD()"
              >close</button
            >
          </template>
        </date-picker>
      </div>
    </div>
    <div
      v-if="choosenDate"
      class="flex justify-center"
    >
      <button
        class="font-semibold text-white bg-foreground hover:text-lgreen text-center text-sm"
        type="button"
        @click="modifyDate()"
      >
        modify date
      </button>
    </div>
    <div
      v-else
      class="flex justify-center"
    >
      <button
        class="font-semibold text-white bg-foreground hover:text-lgreen text-center text-sm"
        type="button"
        @click="selectDate()"
      >
        select date
      </button>
    </div>

    <div
      v-if="choosenDate"
      class="flex flex-row items-stretch justify-between p-5 pt-2 pb-0"
    >
      <label
        for="Times"
        class="justify-self-center self-start text-sm text-white w-10 pt-9"
        >times</label
      >
      <div
        class="flex flex-col justify-self-center"
        id="Times"
      >
        <date-picker
          class="m-2 mr-4"
          type="time"
          value-type="format"
          format="HH:mm"
          v-model:value="sTime"
          v-model:open="openST"
          placeholder="select start time"
          :clearable="true"
          :disabled="choosenDateTimes"
        >
          <template #footer>
            <button
              class="mx-btn mx-btn-text"
              @click="closeST()"
              >close</button
            >
          </template>
        </date-picker>
        <date-picker
          class="m-2 mr-4"
          type="time"
          value-type="format"
          format="HH:mm"
          v-model:value="eTime"
          v-model:open="openET"
          placeholder="select end time"
          :clearable="true"
          :disabled="choosenDateTimes"
        >
          <template #footer>
            <button
              class="mx-btn mx-btn-text"
              @click="closeET()"
              >close</button
            >
          </template>
        </date-picker>
      </div>
    </div>
    <div
      v-if="choosenDate && !choosenDateTimes"
      class="flex justify-center"
    >
      <button
        class="font-semibold text-white bg-foreground hover:text-lgreen text-center text-sm"
        type="button"
        @click="selectTimes()"
      >
        select times
      </button>
    </div>
    <div
      v-if="choosenDateTimes"
      class="flex justify-center"
    >
      <button
        class="font-semibold text-white bg-foreground hover:text-lgreen text-center text-sm"
        type="button"
        @click="modifyTimes()"
      >
        modify times
      </button>
    </div>

    <div
      v-if="choosenDateTimes"
      class="flex justify-end"
    >
      <button
        class="font-semibold text-white border-b border-lgreen bg-foreground hover:bg-lgreen rounded-lg p-2 text-center text-sm"
        type="button"
        @click="
          $emit('changeFromTo', sDateTime, eDateTime);
          sendData();
        "
      >
        confirm
      </button>
    </div>
  </div>
</template>
