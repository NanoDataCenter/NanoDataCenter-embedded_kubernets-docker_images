description_array = [];
function main_form_start() {
    hide_all_sections();
    show_section("main_form");
}
function main_form_init() {
    valve_group_description_map = [];
    for (var i = 0; i < valve_group_names.length; i++) {
        var temp = valve_io[valve_group_names[i]];
        valve_group_description_map.push(valve_group_names[i] + ":" + temp["description"]);
    }
    //console.log("valve_group_names",valve_group_names)
    // console.log("valve_group_description_map",valve_group_description_map)
    jquery_populate_select('#valve_group', valve_group_names, valve_group_description_map, valve_group_change);
    var valve_group_id = valve_group_names[0];
    valves_index = make_valves(valve_io[valve_group_id]);
    valve_choice = valve_group_names[0];
    description_array_a = make_description_array(valve_choice);
    jquery_populate_select("#valve_id", valves_index, description_array_a, valve_change);
}
function make_description_array(valve_choice) {
    return_value = [];
    return_value.push("select valve");
    //console.log('valve_choice',valve_choice)
    //console.log("temp", valve_io[valve_choice]["valve_descriptions"])
    temp = valve_io[valve_choice]["valve_descriptions"];
    //console.log("temp",temp)
    for (i = 0; i < temp.length; i++) {
        return_value.push(i + ":" + temp[i]);
    }
    return return_value;
}
function make_valves(input) {
    temp = input["valve_descriptions"];
    return_value = [];
    return_value.push(0);
    for (i = 1; i < temp.length + 1; i++) {
        return_value.push(i);
    }
    return return_value;
}
function valve_group_change(event, ui) {
    var valve_choice = $("#valve_group").val();
    var valves_index = make_valves(valve_io[valve_choice]);
    description_array_a = make_description_array(valve_choice);
    jquery_populate_select("#valve_id", valves_index, description_array_a, valve_change);
}
function valve_change(event, ui) {
    var index;
    var choice;
    choice = $("#valve_id").val();
    $("#valve_id")[0].selectedIndex = 0;
    if (choice == 0) {
        return;
    }
    var master_controller_id = $("#valve_group")[0].selectedIndex;
    var master_controller_name = valve_group_names[master_controller_id];
    var valve_group_data = valve_io[master_controller_name];
    var stations = valve_group_data["stations"];
    var io = valve_group_data["io"];
    var selected_station = stations[choice];
    var selected_io = io[choice];
    queue_irrigation_direct(selected_station, selected_io);
}
