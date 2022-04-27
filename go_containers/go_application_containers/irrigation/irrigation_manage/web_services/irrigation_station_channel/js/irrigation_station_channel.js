station_list = [];
integer_io_map = {};
function main_form_start() {
    hide_all_sections();
    show_section("main_form");
}
function main_form_init() {
    station_list = Object.keys(io_map);
    console.log(station_list);
    station_list.sort();
    generate_integer_io_map();
    jquery_populate_select('#stations', station_list, station_list, station_change);
    var io_data = integer_io_map[station_list[0]];
    populate_io_list(io_data);
    Time_load_schedule_time("#step_time_time_select", 60);
    $("#step_time_time_select").val('15').change();
}
function Time_load_schedule_time(id, number) {
    load_times = [];
    for (var i = 1; i <= number; i++) {
        load_times.push(i);
    }
    jquery_populate_select(id, load_times, load_times, null);
}
function station_change(event, ui) {
    var station = $("#stations").val();
    var io_list = integer_io_map[station];
    populate_io_list(io_data);
}
function populate_io_list(input_data) {
    jquery_populate_select("#channels", input_data[0], input_data[1], io_change);
}
function io_change(event, ui) {
    var index = $("#channels")[0].selectedIndex;
    io = parseInt($("#channels").val());
    $("#channels")[0].selectedIndex = 0;
    if (index == 0) {
        return;
    }
    var station = $("#stations").val();
    var time = parseInt($("#step_time_time_select").val());
    var message = "station  " + station + " Valve Id " + io;
    queue_irrigation_direct(station, io, time, message);
}
function generate_integer_io_map() {
    for (var i = 0; i < station_list.length; i++) {
        current_station = station_list[i];
        integer_io_map[current_station] = generate_integer_io_map_station(io_map[current_station]);
    }
}
function generate_integer_io_map_station(input_data) {
    var return_value = [];
    var io_channel = [0];
    var description = ["select channel"];
    var temp_dict = generate_integer_channels(input_data);
    var keys = Object.keys(temp_dict);
    keys.sort();
    for (i = 0; i < keys.length; i++) {
        var key = keys[i];
        io_channel.push(key);
        description.push(key + " : " + temp_dict[key]);
    }
    return_value.push(io_channel);
    return_value.push(description);
    return return_value;
}
function generate_integer_channels(input) {
    var return_value = {};
    var keys = Object.keys(input);
    for (i = 0; i < keys.length; i++) {
        var key = keys[i];
        var new_key = parseInt(key);
        return_value[new_key] = input[key];
    }
    return return_value;
}
/*
   {"satellite_1":{"11":"Drip Line along garage","12":"Middle Clover Near Barbecue","13":"Lemon Tree Drip Line near Steps","14":"Barbecue Clover Area","15":"Triangle Pool Area","16":"Pool Fence Area","17":"Dragon Fruit — Fruit Tree Drip Line","18":"Well Clover Area","19":"Well Water Drip Line","20":"Flowers along side walk","21":"Flowers on Opposite Side of Garage","22":"Grass Zone Away From Door","23":"Grass Toward Door","24":"Flowers Toward Garage","25":"Middle Clover Near Well"},"satellite_2":{"1":"Avocado Block6",
  
*/
