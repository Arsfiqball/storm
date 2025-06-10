<template>
  <div class="bg-slate-800 border border-slate-700 rounded-lg p-4 md:p-6">
    <div class="flex items-center mb-4">
      <div class="flex space-x-2">
        <div class="w-3 h-3 rounded-full bg-red-500"></div>
        <div class="w-3 h-3 rounded-full bg-yellow-500"></div>
        <div class="w-3 h-3 rounded-full bg-green-500"></div>
      </div>
    </div>
    <div class="font-mono text-sm md:text-base text-slate-300 overflow-x-auto">
      <p class="pb-2"><span class="text-indigo-400">$</span> storm init-config</p>
      <p class="text-slate-500 pb-2">Configuration file created at ./.storm.yaml</p>
      <p class="pb-2"><span class="text-indigo-400">$</span> storm isolate up --expose --web</p>
      <p class="text-slate-500 pb-2">Starting isolated environment with web service...</p>
      <p class="pb-2"><span class="text-indigo-400">$</span> docker ps</p>
      <p class="text-slate-500">Web frontend running on :5173, backend API on :3000</p>
    </div>
  </div>
</template>
