<template>
  <siderbar class="bg-base-200 h-screen">
    <div class="flex items-center">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-15 w-15 p-3 basis-1/4"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
        />
      </svg>
      <h2 class="basis-3/4 h-15 w-15 p-3 text-center">Backup Tool</h2>
    </div>
    <ul class="menu w-56">
      <li>
        <a class="flex flex-row h-14">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 basis-1/4"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
            />
          </svg>
          <span class="basis-3/4 text-center">备份文件列表</span>
        </a>
      </li>
      <li>
        <a class="flex flex-row h-14" href="/">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 basis-1/4"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
          <span class="basis-3/4 text-center"> Item 2 </span>
        </a>
      </li>
      <li>
        <a class="flex flex-row h-14">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 basis-1/4"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            />
          </svg>
          <span class="basis-3/4 text-center"> Item3</span>
        </a>
      </li>
    </ul>
  </siderbar>
</template>
<script setup lang="ts"></script>
<style lang="">

</style>
