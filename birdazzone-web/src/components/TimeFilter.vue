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
  },
});

</script>

<template>
  <button class="font-semibold text-white border border-lgreen bg-foreground hover:bg-lgreen rounded-lg px-4 py-2 text-center inline-flex items-center" type="button" @click="openCloseFunction()">
    <svg class="mr-2 w-4 h-4" aria-hidden="true" fill="none" stroke="white" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
    </svg>
    filters
  </button>

  <div v-if="openClose" class="z-10 bg-foreground shadow font-semibold text-md rounded-lg m-4 p-4 place-self-center">    
    <div class="flex flex-col" aria-labelledby="dropdownDividerButton">
      <label for="startDateTime" class="justify-self-start text-white">Start</label>
      <div class="flex flex-col justify-self-start" id="startDateTime">
        <date-picker
          class = "m-2"
          v-model:value = "startDate"
          v-model:open="openSD"
          :clearable = "true"
          type = "date"
          placeholder="select start date">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeSD()">close</button>
            </template>
        </date-picker>
        <date-picker
          class = "m-2"
          v-model:value = "startTime" 
          v-model:open="openST"
          :clearable = "true"
          type = "time"
          placeholder="select start time">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeST()">close</button>
            </template>
        </date-picker>
      </div>
    </div>

    <div class="flex flex-col">
      <label for="startDateTime" class="justify-self-start text-white">End</label>
      <div class="flex flex-col justify-self-start" id="startDateTime">
        <date-picker
          class = "m-2"
          v-model:value = "endDate" 
          v-model:open="openED"
          :clearable = "true"
          type = "date"
          placeholder="select end date">
          <template #footer>
            <button class="mx-btn mx-btn-text" @click="closeED()">close</button>
          </template>
        </date-picker>
        <date-picker
          class = "m-2"
          v-model:value = "endTime" 
          v-model:open="openET"
          :clearable = "true"
          type = "time"
          placeholder="select end time">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeET()">close</button>
            </template>
        </date-picker>
      </div>
    </div>

  </div>
</template>
