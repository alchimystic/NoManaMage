local enable_actions = {
	[true] = "enable",
	[false] = "disable",
}

local function get_enable_action(value, defoult)

	local bool = defoult
	if type(value) == 'boolean' then
		bool = value
	end
	return enable_actions[bool]
end

function acquire_input_focus()
	msg.post(".", "acquire_input_focus")
end

function triger_level_load(level)
	msg.post('game:/level_selector#loader', "load", { level = level })
end

function sprite_anim(id)
	msg.post('#sprite', 'play_animation', {id = id})
end

function sprite_enable(value)
	local action = get_enable_action(value, true)
	msg.post('#sprite', action)
end
