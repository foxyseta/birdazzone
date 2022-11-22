<script lang=ts>
import { ref } from 'vue';
import { defineComponent } from 'vue';
import DatePicker from 'vue-datepicker-next';
import 'vue-datepicker-next/index.css';

export default defineComponent({
  components: { DatePicker },
  data() {
    return {
      startDate: null,
      startTime: null,
      endDate: null,
      endTime: null,

      openClose: false,

      openSD: false,
      openST: false,
      openED: false,
      openET: false,
    };
  },
  methods: {
    openCloseFunction() {
      this.openClose = !this.openClose;
    },
    sendData() {
      this.openClose = !this.openClose;
    },
    closeSD() {
      this.openSD = false;
    },
    closeST() {
      this.openST = false;
    },
    closeED() {
      this.openED = false;
    },
    closeET() {
      this.openET = false;
    },
    defaultDate() {
      return new Date();
    },
    disabledAfterToday(date) {
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      return date > today;
    },
    disabledAfterTodayBeforeStartDate(date) {
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      return date > today || date < this.startDate;
    },
    disabledTime(time) {
      
    },
  },
});

</script>

<template>
  <div class="flex justify-end">
    <button class="font-semibold text-white border border-lgreen bg-foreground hover:bg-lgreen rounded-lg px-4 py-2 text-center inline-flex items-center" type="button" @click="openCloseFunction()">
      <svg class="mr-2 w-4 h-4" alt="arrow down" aria-hidden="true" fill="none" stroke="white" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
      </svg>
      filters
    </button>
  </div>

  <div v-if="openClose" class="z-10 bg-foreground shadow font-semibold text-md rounded-lg m-4 mr-0 p-4 place-self-end">    
    <div class="flex flex-row items-stretch" aria-labelledby="dropdownDividerButton">
      <label for="startDateTime" class="justify-self-start self-center text-white w-9">start</label>
      <div class="flex flex-row justify-self-start" id="startDateTime">
        <date-picker
          class = "m-2 mr-1"
          type="date"
          v-model:value="startDate"
          v-model:open="openSD"
          placeholder="select start date"
          :clearable="true"
          :disabled-date="disabledAfterToday">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeSD()">close</button>
            </template>
        </date-picker>
        <date-picker
          class = "m-2 ml-1 mr-0"
          type="time"
          format="HH:mm"
          v-model:value="startTime" 
          v-model:open="openST"
          placeholder="select start time"
          :clearable="true">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeST()">close</button>
            </template>
        </date-picker>
      </div>
    </div>

    <div class="flex flex-row items-stretch">
      <label for="startDateTime" class="justify-self-start self-center text-white w-9">end</label>
      <div class="flex flex-row justify-self-start" id="startDateTime">
        <date-picker
          class = "m-2 mr-1"
          type="date"
          v-model:value="endDate" 
          v-model:open="openED"
          placeholder="select end date"
          :clearable="true"
          :disabled="startDate == null"
          :disabled-date="disabledAfterTodayBeforeStartDate">
          <template #footer>
            <button class="mx-btn mx-btn-text" @click="closeED()">close</button>
          </template>
        </date-picker>
        <date-picker
          class = "m-2 ml-1 mr-0"
          type="time"
          format="HH:mm"
          v-model:value="endTime" 
          v-model:open="openET"
          placeholder="select end time"
          :clearable="true"
          :disabled="startTime == null"
          :disabled-time="disabledTime">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeET()">close</button>
            </template>
        </date-picker>
      </div>
    </div>

    <div class="flex justify-end mt-3">
      <button
        class="font-semibold text-white border border-lgreen bg-foreground hover:bg-lgreen rounded-lg p-2 text-center text-sm"
        type="button" @click="sendData()">
        confirm
      </button>
    </div>

  </div>
</template>
