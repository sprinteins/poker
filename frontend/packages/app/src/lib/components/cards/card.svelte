<!-- https://letsbuildui.dev/articles/a-3d-hover-effect-using-css-transforms -->
<script lang="ts">
  import { CardBacks } from "$lib/card-backs";


    
    // this is just a hack, because the animation is broken 
    // in the card back dialog
    export let disableHoverAnimation = false
    export let label = "" 
    export let bgImagePath = ""
    const THRESHOLD = 10;
    
    let card;
    const defaultCardBack = CardBacks.logo


    function handleHover(e) {
        if(disableHoverAnimation) { return }

        const { pageX, pageY, currentTarget } = e;
        const { clientWidth, clientHeight, offsetLeft, offsetTop } = currentTarget;
        
        const horizontal = (pageX - offsetLeft) / clientWidth;
        const vertical = (pageY - offsetTop) / clientHeight;

        const rotateX = (THRESHOLD / 2 - horizontal * THRESHOLD).toFixed(2);
        const rotateY = (vertical * THRESHOLD - THRESHOLD / 2).toFixed(2);

        

        // 

        card.style.transform = `
            perspective(${clientWidth}px) 
            rotateX(${rotateY}deg) 
            rotateY(${rotateX}deg) 
            scale3d(1, 1, 1)
            translateZ(0.5rem)
        `;
    }

    function resetStyles(e) {
        card.style.transform = `perspective(${e.currentTarget.clientWidth}px) rotateX(0deg) rotateY(0deg)`;
    }
</script>

<div 
    style={`background-image:url('${bgImagePath}')`}
    class="card" 
    class:defaultbg={bgImagePath===defaultCardBack}
    test-id={`card-${label}`}
    bind:this={card} 
    on:mousemove={handleHover} 
    on:mouseleave={resetStyles}
    on:click
>
    <div class="content">
        {label}
    </div>
</div>

<style>

    .card {
        --ratio: 5/3;
        --width: 8rem;
        border: gray thin solid;
        width:  var(--width);
        height: calc( var(--width) * (var(--ratio)) );
        
        display:     grid;
        place-items: center;
        
        /* background:    black; */
        color:         white;
        cursor:        pointer; 
        border-radius: 4px;
        
        background:          var(--color-gray-dark);
        background-size:     100%;
        background-repeat:   no-repeat;
        background-position: center;    
    }

    .card {
        position:        relative;
        /* transition:      transform 0.01s ease; */
        transform-style: preserve-3d;
        will-change:     transform;
    }

    .card.defaultbg{
        background-size: 5rem;
    }

    .card:hover{
        transition:         transform 0.1s;
        -webkit-box-shadow: 0px 0px 21px -2px #000000; 
        box-shadow:         0px 0px 21px -2px #000000;
    }

    .card:hover .content {
        transform:   translateZ(12px);
        text-shadow: 0px 0px 8px #FFFFFF;
    }

    .content {
        font-size:   clamp(1.5rem, 4vw, 3.25rem);
        font-family: "Helvetica Neue";
        font-weight: bold;
        position:    relative;
        z-index:     1;
        max-width:   100%;
        padding:     0 0.5rem;
        line-height: 1;
        text-align:  center;
        /* transition:  transform 0.1s ease; */
        transition:  text-shadow;
        /* text-shadow: 0px 0px 8px #FFFFFF; */
        text-transform: uppercase;
        overflow-wrap: anywhere;
    }

    .card::before {
        content:    "";
        background: rgba(0, 0, 0, 0.4);
        position:   absolute;
        height:     100%;
        width:      100%;
        left:       0;
        right:      0;
        top:        0;
        bottom:     0;
    }

</style>