description_array = [];
function valve_group_components_start() {
    initialize_direct_io_control();
    hide_all_sections();
    show_section("valve_group_components");
}
function valve_group_components_init() {
    valve_group_description_map = [];
    for (var i = 0; i < valve_group_names.length; i++) {
        var temp = valve_io[valve_group_names[i]];
        valve_group_description_map.push(valve_group_names[i] + ":" + temp["description"]);
    }
    attach_button_handler("#valve_group_cancel_id", station_channel_cancel_id);
    //console.log("valve_group_names",valve_group_names)
    // console.log("valve_group_description_map",valve_group_description_map)
    jquery_populate_select('#valve_group', valve_group_names, valve_group_description_map, valve_group_change);
    var valve_group_id = valve_group_names[0];
    valves_index = make_valves(valve_io[valve_group_id]);
    valve_choice = valve_group_names[0];
    description_array_a = make_description_array(valve_choice);
    jquery_populate_select("#valve_id", valves_index, description_array_a, valve_change);
    Time_load_schedule_time("#valve_step_time_time_select", 60);
    $("#valve_step_time_time_select").val('15').change();
}
function valve_group_cancel_id() {
    start_section("main_form");
}
function Time_load_schedule_time(id, number) {
    load_times = [];
    for (var i = 1; i <= number; i++) {
        load_times.push(i);
    }
    jquery_populate_select(id, load_times, load_times, null);
}
function make_description_array(valve_choice) {
    return_value = [];
    return_value.push("select valve");
    //console.log('valve_choice',valve_choice)
    //console.log("temp", valve_io[valve_choice]["valve_descriptions"])
    temp = valve_io[valve_choice]["valve_descriptions"];
    //console.log("temp",temp)
    for (i = 0; i < temp.length; i++) {
        return_value.push((i + 1) + ":" + temp[i]);
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
/*
 * relation "time_idx" already exists (SQLSTATE 42P07) CREATE INDEX time_idx ON T8878891975870246288(Time);
valve io  {"valve group 1":{"description":"xxxxxxxxxxxxxxx","io":[20,19,14,18,15,17,16],"stations":["station_1","station_1","station_1","station_1","station_1","station_1","station_1"],"valve_descriptions":["Flowers along side walk","Well Water Drip Line","Barbecue Clover Area","Well Clover Area","Triangle Pool Area","Dragon Fruit — Fruit Tree Drip Line","Pool Fence Area"]},"valve group 10":{"description":"xxxxxxxxxxxxxxx","io":[1,2,3,4,5,6,7,8],"stations":["station_4","station_4","station_4","station_4","station_4","station_4","station_4","station_4"],"valve_descriptions":["???????????
*/
function valve_change(event, ui) {
    var index;
    var valve_group_id;
    if ($("#valve_id")[0].selectedIndex == 0) {
        return;
    }
    valve_group_id = parseInt($("#valve_id").val());
    var valve_group = $("#valve_group").val();
    var item = valve_io[valve_group];
    console.log("item", item);
    valve_group_id = valve_group_id - 1;
    var station = item["stations"][valve_group_id];
    var io = item["io"][valve_group_id];
    var time = parseInt($("#valve_step_time_time_select").val());
    console.log("valve data ", station, io, time);
    queue_irrigation_direct(station, io, time, "Queue Valve Group " + valve_group + "  Valve Number " + (valve_group_id + 1));
    $("#valve_id")[0].selectedIndex = 0;
    /*
     let valve_group_data               = valve_io[master_controller_name]
   
     let stations                                = valve_group_data["stations"]
     let io                                           = valve_group_data["io"]
     let selected_station                = stations[choice]
     let selected_io                          = io[choice]
     let time                                     = parseInt($("#valve_step_time_time_select").val())
     let message = "Queue Valve Group  "+master_controller_name +" Valve Id "+choice
    queue_irrigation_direct(selected_station ,selected_io,time,message)
    */
}
