<script>
	import { goto } from "$app/navigation";
    import HomeIcon from "$lib/components/atom/icon/HomeIcon.svelte";
	import SideBarIcon from "$lib/components/atom/SideBarIcon.svelte";
import "../app.css";
</script>
  
<div class="bg-white h-full w-[5%] absolute">
<SideBarIcon onClick={() => goto('/')}><HomeIcon></HomeIcon></SideBarIcon>
</div>

  <div class="bg-[#55868C] text-center p-5 h-screen flex flex-col items-center">
      <slot />
  </div>