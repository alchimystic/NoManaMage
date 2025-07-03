
function triger_level_load(level)
	msg.post('game:/level_selector#loader', "load", { level = level })
end
