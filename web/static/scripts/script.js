(function() {
  const updateBtns = document.querySelectorAll('[data-action="update"]');

  updateBtns.forEach(btn => {
    btn.addEventListener('click', () => {
      const parent = btn.closest('[role="listitem"]');
      const updateForm = parent.querySelector('[data-form="update"]');
      const input = updateForm.querySelector('input[type="text"]');
      const currentValue = input.value;
      const cancelBtn = parent.querySelector('[data-action="cancel"]');

      updateForm.style.display = 'block';
      
      cancelBtn.addEventListener('click', () => {
        updateForm.style.display = 'none';
        input.value = currentValue;
      });
    });
  })
})();