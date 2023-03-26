<template>
  <div class="flex flex-col min-h-screen">
    <main class="w-[100%] h-[100%]">
        <form ref="form" method="POST" action="/process-image" enctype="multipart/form-data" class="container mx-auto flex flex-row gap-4 justify-center">
          <div class="w-1/3 max-w-12 flex flex-col gap-4 justify-center min-h-screen">
            <fieldset>
              <legend class="block text-sm font-medium leading-6 text-gray-900">Output size</legend>
              <div class="mt-2 -space-y-px rounded-md bg-white shadow-sm">
                <div class="flex -space-x-px">
                  <div class="w-1/2 min-w-0 flex-1">
                    <label for="width" class="sr-only">Width</label>
                    <input type="text" name="width" value="1920" id="width" class="text-center relative block w-full rounded-l-md border-0 bg-transparent py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" placeholder="1920" />
                  </div>
                  <div class="min-w-0 flex-1">
                    <label for="height" class="sr-only">Height</label>
                    <input type="text" name="height" id="height" value="1080" class="text-center relative block w-full rounded-r-md border-0 bg-transparent py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" placeholder="1080" />
                  </div>
                </div>
              </div>
            </fieldset>

            <fieldset class="flex align-baseline justify-between">
              <span class="font-medium text-sm">Do you want black background?</span>
              <Switch v-model="wantBlackBackground" :class="[wantBlackBackground ? 'bg-indigo-600' : 'bg-gray-200', 'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-600 focus:ring-offset-2']">
                <span class="sr-only">Use setting</span>
                <span aria-hidden="true" :class="[wantBlackBackground ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out']" />
              </Switch>
              <input type="hidden" id="black_background" name="black_background" :value="wantBlackBackground">
            </fieldset>

            <!-- <FileUpload @file-uploaded="handleFile"/> -->

            <fieldset>
              <label class="block mb-2 text-sm font-medium text-gray-900" for="file_input">Original file</label>
              <input class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 focus:outline-none" id="file_input" type="file" accept="image/jpeg" name="image" required>
            </fieldset>

            <button class="rounded-md bg-indigo-600 py-2.5 px-3.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Submit</button>
          </div>
        </form>
    </main>
    <footer class="mt-auto mb-2 container mx-auto px-4 flex justify-between flex-wrap">
      <span>Image Scratch &mdash; 2023</span>
      <span>theeldarka</span>
    </footer>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { Switch } from '@headlessui/vue'
import FileUpload from './components/FileUpload.vue'
import { onMounted } from 'vue'
import { initFlowbite } from 'flowbite'

const wantBlackBackground = ref(true)

onMounted(() => {
    initFlowbite();
})
</script>

<script>

export default {
  components: {
    Switch,
    FileUpload
  },
  methods: {
    handleFile(files) {
      console.log(files)

      this.$refs.form.submit()
    }
  }
}

</script>