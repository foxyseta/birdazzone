<script lang=ts>
import { checkCompatEnabled } from '@vue/compiler-core';
import { ref } from 'vue';
import { defineComponent } from 'vue';
import DatePicker from 'vue-datepicker-next';
import 'vue-datepicker-next/index.css';
import ErrorWidget from './ErrorWidget.vue';


export default defineComponent({
  components: { DatePicker },
  props: {
    from: {
      type: String,
      default: '',
    },
    to: {
      type: String,
      default: '',
    }
  },

  data() {
    return {
      dates: null,          /** to save dates values */
      sDate: String(null),    /** to save start date */
      eDate: String(null),    /** to save end date */
      choosenDates: false,  /** to verify if dates has been chosen */
      equalDates: ref(false),    /** to verify if dates are equals */
      diffDates: ref(false),     /** to verify if dates are equals */
      openD: false,         /** to close dates popup */

      sTime: String(null),  /** to same start time value */
      openST: false,          /** to close start time popup */
      eTime: String(null),  /** to same end time value */
      openET: false,          /** to close end time popup */

      DatesTimes: false,
      openClose: false,   /** to open/close filters card */
    };
  },

  methods: {
    /** GENERAL */
    openCloseFunction() {               /** to open/close filters card */
      this.openClose = !this.openClose;
    },

    /** DATES */
    closeD() {                          /** to close dates popup */
      this.openD = false;
    },
    selectDates(){                      /** to confirm selected dates */
      if (this.dates != null) {              // entered dates
        this.sDate = this.dates[0];
        this.eDate = this.dates[1];
      }   // ELSE -> didnt enter dates -> by default: today

      if(this.sDate === this.eDate){        // equal dates
        this.choosenDates = true;
        this.equalDates = true;
      }
      else {                                // different dates
        this.sDate = this.sDate + "T00:00:00.000Z";
        this.eDate = this.eDate + "T00:00:00.000Z";
        this.equalDates = false;
        this.diffDates = true;
      }

      console.log("start date: " + this.sDate);
      console.log("end date: " + this.eDate);
    },

    /** TIMES */
    closeST() {                          /** to close start time popup */
      this.openST = false;
    },
    closeET() {                          /** to close end time popup */
      this.openET = false;
    },
    disabledAfterToday(date) {          /** to disabilitate days after today in dates popup */
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      return date > today;
    },
    selectTimes(){
      if(this.sDate == this.eDate && this.sTime != null && this.eTime != null){                 /* entered same valid date: check on entered times */ 
        if(this.sTime.substring(0, 2) === this.eTime.substring(0, 2))  {            // sTime HH == eTime HH (same hours)
          if(this.sTime.substring(3) > this.eTime.substring(3)){                      // sTime mm > eTime mm -> switch of the two dates
            let box = this.sTime;
            this.sTime = this.eTime;
            this.eTime = box;
          }
        }
        else{                                                                      // sTime HH != eTime HH (different hours)
          if(this.sTime.substring(0, 2) > this.eTime.substring(0, 2)){                  // sTime HH > eTime HH -> switch of the two
            let box = this.sTime;
            this.sTime = this.eTime;
            this.eTime = box;
          }
        }
        this.sDate = this.sDate + "T" + this.sTime + ":00.000Z";
        this.eDate = this.eDate + "T" + this.eTime + ":00.000Z";

        this.DatesTimes = true;
      }
    },

    /** GENERAL */
    sendData() {                        /** confirm button */
      this.openCloseFunction();
      console.log("from: " + this.from);
      console.log("to: " + this.to);

      console.log("start date: " + this.sDate);
      console.log("end date: " + this.eDate);
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
      <label for="Dates" class="justify-self-start self-center text-white w-9">dates</label>
      <div class="flex justify-self-start" id="Dates">
        <date-picker 
          class = "m-2 ml-3"
          range
          type="date"
          v-model:value="dates"
          v-model:open="openD"
          value-type="format"
          format="YYYY-MM-DD"
          placeholder="select start and end date"
          :clearable="true"
          :disabled-date="disabledAfterToday"
          @change="!equalDates">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeD()">close</button>
            </template>
        </date-picker>
      </div>
    </div>
    <div class="flex justify-center">
      <button class="font-semibold text-white bg-foreground hover:text-lgreen text-center text-sm" type="button"
        @click="selectDates()">
        select dates
      </button>
    </div>

    <div v-if="equalDates" class="flex flex-row items-stretch">
      <label for="Times" class="justify-self-start self-center text-white w-9">times</label>
      <div class="flex flex-row justify-self-start" id="Times">
        <date-picker 
          class = "m-2 ml-3 mr-1"
          type="time"
          value-type="format"
          format="HH:mm"
          v-model:value="sTime"
          v-model:open="openST"
          placeholder="select start time"
          :clearable="true">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeST()">close</button>
            </template>
        </date-picker>
        <date-picker
          class = "m-2 mr-0"
          type="time"
          value-type="format"
          format="HH:mm"
          v-model:value="eTime"
          v-model:open="openET"
          placeholder="select end time"
          :clearable="true"
          :disabled="sTime == null">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeET()">close</button>
            </template>
        </date-picker>
      </div>
      <div class="flex justify-center">
        <button class="font-semibold text-white bg-foreground hover:text-lgreen text-center text-sm" type="button" @click="selectTimes()">
          select times
        </button>
      </div>
    </div>

    <div v-if="diffDates || DatesTimes" class="flex justify-end mt-3">
      <button
        class="font-semibold text-white border border-lgreen bg-foreground hover:bg-lgreen rounded-lg p-2 text-center text-sm"
        type="button" @click="$emit('changeFrom', sDate); $emit('changeTo', eDate); sendData()">
        confirm
      </button>
    </div>

  </div>
</template>
