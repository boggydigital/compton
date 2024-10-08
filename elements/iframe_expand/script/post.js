document.addEventListener("DOMContentLoaded", () => {
    const ro = new ResizeObserver(entries => {
      for (let e of entries) {
        let msg = {
            context: ctx,
            height: e.contentRect.height,
        }
        if (window.parent) {
            window.parent.postMessage(JSON.stringify(msg), "*");
        }
      }
    });
    const ctx = document.body.id
    ro.observe(document.getElementById(ctx));
});
