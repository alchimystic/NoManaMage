
function acquire_input_focus()
	msg.post(".", "acquire_input_focus")
end

function triger_level_load(level)
	msg.post('game:/level_selector#loader', "load", { level = level })
end

